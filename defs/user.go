package defs

type ServiceUser struct {
	Mail                  string `json:"mail"`
	Active                bool
	LoginToken            string `json:"login_token"`
	MailVeryficationToken string `json:"mail_veryfication_token"`
}
