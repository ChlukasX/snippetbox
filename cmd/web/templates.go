package main

import "snippetbox.chlukas.github/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}