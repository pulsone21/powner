package charts

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/pulsone21/powner/internal/entities"
)

type RadarChart struct {
	Labels   []string  `json:"labels"`
	Datasets []Dataset `json:"datasets"`
}

type Dataset struct {
	Label                     string `json:"label"`
	Data                      []int  `json:"data"`
	Fill                      bool   `json:"fill"`
	BackgroundColor           string `json:"backgroundColor"`
	BorderColor               string `json:"borderColor"`
	PointBackgroundColor      string `json:"pointBackgroundColor"`
	PointBorderColor          string `json:"pointBorderColor"`
	PointHoverBackgroundColor string `json:"pointHoverBackgroundColor"`
	PointHoverBorderColor     string `json:"pointHoverBorderColor"`
}

func maxData(datapoints int) Dataset {
	data := make([]int, datapoints)
	for i := 0; i < datapoints; i++ {
		data[i] = 5
	}
	return Dataset{
		Label:                     "Maximum",
		Data:                      data,
		Fill:                      true,
		BackgroundColor:           "rgba(255, 179, 203, 0.5)",
		BorderColor:               "#df677b",
		PointBackgroundColor:      "rgba(223, 103, 123, 0.5)",
		PointHoverBorderColor:     "#efb3bd",
		PointBorderColor:          "#df677b",
		PointHoverBackgroundColor: "rgba(239, 179, 189, 0.5)",
	}
}

func newRadarChart(t entities.Team) RadarChart {
	slog.Info("Creating new RadarChart")
	log.Println(t)

	labels := []string{}
	avg_data := Dataset{
		Label:                     "Team Average",
		Data:                      []int{},
		Fill:                      true,
		BackgroundColor:           "rgba(166, 227, 161, 0.5)",
		BorderColor:               "#67cf5e",
		PointBackgroundColor:      "rgba(103, 207, 94, 0.5)",
		PointHoverBorderColor:     "#cbefc8",
		PointBorderColor:          "#67cf5e",
		PointHoverBackgroundColor: "rgba(203, 239, 200, 0.5)",
	}

	max_data := Dataset{
		Label:                     "Team Maximum",
		Data:                      []int{},
		Fill:                      true,
		BackgroundColor:           "rgba(255, 179, 203, 0.5)",
		BorderColor:               "#df677b",
		PointBackgroundColor:      "rgba(223, 103, 123, 0.5)",
		PointHoverBorderColor:     "#efb3bd",
		PointBorderColor:          "#df677b",
		PointHoverBackgroundColor: "rgba(239, 179, 189, 0.5)",
	}

	for _, s := range t.Skills {
		labels = append(labels, s.Name)
		max_team := 0
		for _, m := range t.Members {
			for _, sr := range m.Skills {
				if sr.Skill.ID == s.ID && max_team < sr.Rating {
					max_team = sr.Rating
				}
			}
		}

		slog.Debug(fmt.Sprintf("Trying getting avg out of maxTeam=%v and lenMembers=%v\n", max_team, len(t.Members)))
		avg_data.Data = append(avg_data.Data, max_team/len(t.Members))
		max_data.Data = append(max_data.Data, max_team)

	}

	return RadarChart{
		Labels:   labels,
		Datasets: []Dataset{max_data, avg_data},
	}
}
