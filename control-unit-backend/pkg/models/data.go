package models

import (
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
	datas        []Data
	historyDatas []HistoryData
	datasBuffer  []Data
	mu           sync.Mutex
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
	if len(s.historyDatas) == 0 {
		return HistoryData{-1, -1, -1, ""}
	}
	return s.historyDatas[len(s.historyDatas)-1]
}

func (s *Sampler) GetHistoryDatas() []HistoryData {
	return s.historyDatas
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

			s.historyDatas = append(s.historyDatas, HistoryData{
				Avg:  avg,
				Min:  min,
				Max:  max,
				Date: timestamp,
			})
			s.datasBuffer = []Data{}
			s.mu.Unlock()
		}
	}()
}

var DataSampler = Sampler{}
