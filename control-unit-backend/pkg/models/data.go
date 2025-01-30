package models

import (
	"sync"
	"time"
)

const N_TEMP_MEASUREMENTS int = 50 // we keep track of last N measurements
const PERIOD_MIN int = 5           // every period we produce a avg-min-max sample

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
	Datas        []Data
	HistoryDatas []HistoryData
	DatasBuffer  []Data
	mu           sync.Mutex
}

func (s *Sampler) AddData(temp float32, date string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var data = Data{temp, date}
	if len(s.Datas) < N_TEMP_MEASUREMENTS {
		s.Datas = append(s.Datas, data)
	} else {
		s.Datas = s.Datas[1:] // remove the first element
		s.Datas = append(s.Datas, data)
	}
	s.DatasBuffer = append(s.DatasBuffer, data)
}

func (s *Sampler) GetLastData() Data {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.Datas) == 0 {
		return Data{0, ""}
	}
	return s.Datas[len(s.Datas)-1]
}

// Go subroutine that samples datas creating avg, min and max every PERIOD_MIN minutes
func (s *Sampler) StartSampling() {
	ticker := time.NewTicker(time.Duration(PERIOD_MIN) * time.Minute)
	go func() {
		for range ticker.C {
			s.mu.Lock()
			if len(s.DatasBuffer) == 0 { // Sampler is empty
				s.mu.Unlock()
				continue
			}

			var sum, min, max float32
			min = s.DatasBuffer[0].Temp
			max = s.DatasBuffer[0].Temp

			for _, d := range s.DatasBuffer {
				sum += d.Temp
				if d.Temp < min {
					min = d.Temp
				}
				if d.Temp > max {
					max = d.Temp
				}
			}
			avg := sum / float32(len(s.DatasBuffer))
			timestamp := time.Now().Format(time.RFC3339)

			s.HistoryDatas = append(s.HistoryDatas, HistoryData{
				Avg:  avg,
				Min:  min,
				Max:  max,
				Date: timestamp,
			})
			s.DatasBuffer = []Data{}
			s.mu.Unlock()
		}
	}()
}
