package controllers

import (
	"net/http"

	"github.com/progsamdev/coursescalhoun/views"
)

func StaticHandler(tpl *views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl *views.Template) http.HandlerFunc {

	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "Q: How can I contact you?",
			Answer:   `A: You can reach me via email at <a href="mailto:samuel.msbr@gmail.com">samuel.msbr@gmail.com" </a>`,
		},
		{
			Question: "Q: What services do you provide?",
			Answer:   "A: Currently, I do not provide specific services through this platform. If you have any inquiries, feel free to contact me via email.",
		},
		{
			Question: "Q: Can I collaborate with you on a project?",
			Answer:   "A: I'm open to collaboration opportunities. Please reach out via email to discuss further details.",
		},
		{
			Question: "Q: How long does it take to receive a response? ",
			Answer:   "A: I strive to respond to emails as promptly as possible. However, response times may vary based on workload and other commitments.",
		},
		{
			Question: "Q: Do you have a social media?",
			Answer:   " A: At the moment, I don't have active social media profiles. Email is the best way to get in touch.",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
