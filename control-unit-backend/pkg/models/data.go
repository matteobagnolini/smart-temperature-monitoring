package models

import (
	"control-unit-backend/pkg/db"
	"sync"
	"time"
)

const N_TEMP_MEASUREMENTS int = 50 // we keep track of last N measurements
const PERIOD_MIN int = 1           // every period we produce a avg-min-max sample

type Data struct {
	Temp float32
	Date string
}

type HistoryData struct {
	Avg  float32
	Min  float32
	Max  float32
	Date string
}

type Sampler struct {
	datas           []Data
	lastHistoryData HistoryData // last history data is cached
	datasBuffer     []Data
	mu              sync.Mutex
}

func (s *Sampler) AddData(temp float32, date string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var data = Data{temp, date}
	if len(s.datas) < N_TEMP_MEASUREMENTS {
		s.datas = append(s.datas, data)
	} else {
		s.datas = s.datas[1:] // remove the first element
		s.datas = append(s.datas, data)
	}
	s.datasBuffer = append(s.datasBuffer, data)
}

func (s *Sampler) GetLastData() Data {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.datas) == 0 {
		return Data{-1, ""}
	}
	return s.datas[len(s.datas)-1]
}

func (s *Sampler) GetDatas() []Data {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.datas) == 0 {
		return []Data{{-1, ""}}
	}
	return s.datas
}

func (s *Sampler) GetLastHistoryData() HistoryData {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.lastHistoryData
}

func (s *Sampler) GetHistoryDatas() []HistoryData {
	s.mu.Lock()
	defer s.mu.Unlock()
	var datas []db.Dbdata = db.GetAllDatas()
	var historyDatas []HistoryData
	for _, data := range datas {
		historyDatas = append(historyDatas, HistoryData{
			data.Avg,
			data.Min,
			data.Max,
			data.Date,
		})
	}
	return historyDatas
}

// Go subroutine that samples datas creating avg, min and max every PERIOD_MIN minutes
func (s *Sampler) StartSampling() {
	ticker := time.NewTicker(time.Duration(PERIOD_MIN) * time.Minute)
	go func() {
		for range ticker.C {
			s.mu.Lock()
			if len(s.datasBuffer) == 0 { // Sampler is empty
				s.mu.Unlock()
				continue
			}

			var sum, min, max float32
			min = s.datasBuffer[0].Temp
			max = s.datasBuffer[0].Temp

			for _, d := range s.datasBuffer {
				sum += d.Temp
				if d.Temp < min {
					min = d.Temp
				}
				if d.Temp > max {
					max = d.Temp
				}
			}
			avg := sum / float32(len(s.datasBuffer))
			timestamp := time.Now().Format(time.ANSIC)

			s.lastHistoryData = HistoryData{
				Avg:  avg,
				Min:  min,
				Max:  max,
				Date: timestamp,
			}
			db.AddData(avg, min, max, timestamp)

			s.datasBuffer = []Data{}
			s.mu.Unlock()
		}
	}()
}

var DataSampler = Sampler{}
