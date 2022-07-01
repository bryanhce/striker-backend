package app

import (
	"backend/cmd/models"
	"fmt"
	"net/http"
	"time"

	sendgrid "github.com/sendgrid/sendgrid-go"
	mail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func sendEmail(email, taskText string) error {
    from := mail.NewEmail("Striker", "coding.cucumbers@gmail.com")
    subject := "Task Deadline"
    to := mail.NewEmail("User", email)
	date := time.Now().Format("01/02/2006")
    plainTextContent := ""
    htmlContent := fmt.Sprintf(`
	<h2>Hello!</h2>
	<span>Just a friendly reminder the deadline for your tasks
	 	<strong>%s</strong> 
	are today, %s!</span>
	<br/>
	<br/>
	<span>Stay productive,<span>
	<br/>
	<span>Team Head in the Clouds<span>`, taskText, date)
    message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	apiKey := "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"
    client := sendgrid.NewSendClient(apiKey)
    res, err := client.Send(message)
    if err != nil {
		fmt.Println(res)
		return err
    } else {
		fmt.Println(res)
        return nil
    }
}

func (a *App) PostReminderEmails(w http.ResponseWriter, r *http.Request) {
	emailList, err := models.GetReminderEmails(a.DB)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	for _, re := range *emailList {
		err = sendEmail(re.Email, re.Description)
		if err != nil {
			fmt.Sprintf("error sending email to %s, error %s", re.Email, err.Error())
			continue
		}
	}
	
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "emails sent"})
}