package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {

	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "How can I contact you?",
			Answer:   `You can reach me via email at <a href="mailto:samuel.msbr@gmail.com">samuel.msbr@gmail.com </a>`,
		},
		{
			Question: "What services do you provide?",
			Answer:   "Currently, I do not provide specific services through this platform. If you have any inquiries, feel free to contact me via email.",
		},
		{
			Question: "Can I collaborate with you on a project?",
			Answer:   "I'm open to collaboration opportunities. Please reach out via email to discuss further details.",
		},
		{
			Question: "How long does it take to receive a response? ",
			Answer:   "I strive to respond to emails as promptly as possible. However, response times may vary based on workload and other commitments.",
		},
		{
			Question: "Do you have a social media?",
			Answer:   "At the moment, I don't have active social media profiles. Email is the best way to get in touch.",
		},
		{
			Question: "Do you have a social media?",
			Answer:   "At the moment, I don't have active social media profiles. Email is the best way to get in touch.",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
