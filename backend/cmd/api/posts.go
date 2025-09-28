package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rayaadhary/social-go/internal/store"
)

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var input store.Post

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body")
	}

	if err := app.services.Posts.CreatePost(r.Context(), &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeSuccess(w, http.StatusCreated, input, "post created successfully")
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	post, err := app.services.Posts.GetPost(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusNotFound, "post not found")
		return
	}

	writeSuccess(w, http.StatusOK, post, "success get post")
}

func (app *application) listPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := app.services.Posts.ListPosts(r.Context(), 10, 0)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to fetch posts")
		return
	}

	writeSuccess(w, http.StatusOK, posts, "success list posts")
}

func (app *application) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var input store.Post
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body")
		return
	}
	input.ID = id

	if err := app.services.Posts.UpdatePost(r.Context(), &input); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeSuccess(w, http.StatusOK, input, "post updated successfully")
}

func (app *application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := app.services.Posts.DeletePost(r.Context(), id); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeSuccess(w, http.StatusOK, nil, "post deleted successfully")
}
