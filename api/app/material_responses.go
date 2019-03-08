// InfoMark - a platform for managing courses with
//            distributing exercise sheets and testing exercise submissions
// Copyright (C) 2019  ComputerGraphics Tuebingen
// Authors: Patrick Wieschollek
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package app

import (
	"net/http"
	"time"

	"github.com/cgtuebingen/infomark-backend/model"
	"github.com/go-chi/render"
)

// MaterialResponse is the response payload for Material management.
type MaterialResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Filename  string    `json:"filename"` // we keep the original name, since it is meaningful
	Kind      int       `json:"kind"`
	PublishAt time.Time `json:"publish_at"`
	LectureAt time.Time `json:"lecture_at"`
}

// newMaterialResponse creates a response from a Material model.
func (rs *MaterialResource) newMaterialResponse(p *model.Material) *MaterialResponse {
	return &MaterialResponse{
		ID:        p.ID,
		Name:      p.Name,
		Filename:  p.Filename,
		Kind:      p.Kind,
		PublishAt: p.PublishAt,
		LectureAt: p.LectureAt,
	}
}

// newMaterialListResponse creates a response from a list of Material models.
func (rs *MaterialResource) newMaterialListResponse(Materials []model.Material) []render.Renderer {
	// https://stackoverflow.com/a/36463641/7443104
	list := []render.Renderer{}
	for k := range Materials {
		list = append(list, rs.newMaterialResponse(&Materials[k]))
	}
	return list
}

// Render post-processes a MaterialResponse.
func (body *MaterialResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}