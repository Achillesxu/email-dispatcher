package dispatcher

import (
	"latest/config"
	"latest/dto"
	"log"

	"gopkg.in/gomail.v2"
)

type MailService struct {
	smtp     string
	user     string
	password string
	port     uint
}

type MailServiceMessage struct {
	From    string
	To      string
	Subject string
	Body    string
}

func (c *MailService) GetMailInstance() *gomail.Dialer {
	return gomail.NewDialer(c.smtp, int(c.port), c.user, c.password)
}

func NewMailService() *MailService {
	return &MailService{
		smtp:     config.GetConfig().MailSmtp,
		user:     config.GetConfig().MailUser,
		password: config.GetConfig().MailPassword,
		port:     config.GetConfig().MailPort,
	}
}

func CodeMessage(d dto.KafkaResponse) *gomail.Message {

	str := codeTemplateHTML(d.Code)

	y := gomail.NewMessage()
	y.SetHeader("From", config.GetConfig().MailUser)
	y.SetHeader("To", d.Email)
	y.SetHeader("Subject", "Greetings from us - Go service")
	y.SetBody("text/html", str)

	return y
}

func ChangeEmailMessage(d dto.KafkaResponse) error {

	str := newEmailTemplateHTML(d.Code)

	y := gomail.NewMessage()
	y.SetHeader("From", config.GetConfig().MailUser)
	y.SetHeader("To", d.Email)
	y.SetHeader("Subject", "Confirm new email - Go service")
	y.SetBody("text/html", str)

	if err := mail.DialAndSend(m); err != nil {
		log.Println(err)
		return err
	}

	return y
}

func LostPassMessage(d dto.KafkaResponse) *gomail.Message {

	y := gomail.NewMessage()
	y.SetHeader("From", config.GetConfig().MailUser)
	y.SetHeader("To", d.Email)
	y.SetHeader("Subject", "Lost password - Go service")
	y.SetBody("text/html", str)

	return y
}

func templateHTML(code string) string {

	str := `
	<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<html xmlns="http://www.w3.org/1999/xhtml">
	<head>
	  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
	  <meta name="viewport" content="width=device-width, initial-scale=1" />
	  <title>Oxygen Welcome</title>
	
	  <style type="text/css">
		/* Take care of image borders and formatting, client hacks */
		img { max-width: 600px; outline: none; text-decoration: none; -ms-interpolation-mode: bicubic;}
		a img { border: none; }
		table { border-collapse: collapse !important;}
		#outlook a { padding:0; }
		.ReadMsgBody { width: 100%; }
		.ExternalClass { width: 100%; }
		.backgroundTable { margin: 0 auto; padding: 0; width: 100% !important; }
		table td { border-collapse: collapse; }
		.ExternalClass * { line-height: 115%; }
		.container-for-gmail-android { min-width: 600px; }
	
	
		/* General styling */
		* {
		  font-family: Helvetica, Arial, sans-serif;
		}
	
		body {
		  -webkit-font-smoothing: antialiased;
		  -webkit-text-size-adjust: none;
		  width: 100% !important;
		  margin: 0 !important;
		  height: 100%;
		  color: #676767;
		}
	
		td {
		  font-family: Helvetica, Arial, sans-serif;
		  font-size: 14px;
		  color: #777777;
		  text-align: center;
		  line-height: 21px;
		}
	
		a {
		  color: #676767;
		  text-decoration: none !important;
		}
	
		.pull-left {
		  text-align: left;
		}
	
		.pull-right {
		  text-align: right;
		}
	
		.header-lg,
		.header-md,
		.header-sm {
		  font-size: 32px;
		  font-weight: 700;
		  line-height: normal;
		  padding: 35px 0 0;
		  color: #4d4d4d;
		}
	
		.header-md {
		  font-size: 24px;
		}
	
		.header-sm {
		  padding: 5px 0;
		  font-size: 18px;
		  line-height: 1.3;
		}
	
		.content-padding {
		  padding: 20px 0 30px;
		}
	
		.mobile-header-padding-right {
		  width: 290px;
		  text-align: right;
		  padding-left: 10px;
		}
	
		.mobile-header-padding-left {
		  width: 290px;
		  text-align: left;
		  padding-left: 10px;
		}
	
		.free-text {
		  width: 100% !important;
		  padding: 10px 60px 0px;
		}
	
		.block-rounded {
		  border-radius: 5px;
		  border: 1px solid #e5e5e5;
		  vertical-align: top;
		}
	
		.button {
		  padding: 30px 0;
		}
	
		.info-block {
		  padding: 0 20px;
		  width: 260px;
		}
	
		.block-rounded {
		  width: 260px;
		}
	
		.info-img {
		  width: 258px;
		  border-radius: 5px 5px 0 0;
		}
	
		.force-width-gmail {
		  min-width:600px;
		  height: 0px !important;
		  line-height: 1px !important;
		  font-size: 1px !important;
		}
	
		.button-width {
		  width: 228px;
		}
	
	  </style>
	
	  <style type="text/css" media="screen">
		@import url(http://fonts.googleapis.com/css?family=Oxygen:400,700);
	  </style>
	
	  <style type="text/css" media="screen">
		@media screen {
		  /* Thanks Outlook 2013! */
		  * {
			font-family: 'Oxygen', 'Helvetica Neue', 'Arial', 'sans-serif' !important;
		  }
		}
	  </style>
	
	  <style type="text/css" media="only screen and (max-width: 480px)">
		/* Mobile styles */
		@media only screen and (max-width: 480px) {
	
		  table[class*="container-for-gmail-android"] {
			min-width: 290px !important;
			width: 100% !important;
		  }
	
		  table[class="w320"] {
			width: 320px !important;
		  }
	
		  img[class="force-width-gmail"] {
			display: none !important;
			width: 0 !important;
			height: 0 !important;
		  }
	
		  a[class="button-width"],
		  a[class="button-mobile"] {
			width: 248px !important;
		  }
	
		  td[class*="mobile-header-padding-left"] {
			width: 160px !important;
			padding-left: 0 !important;
		  }
	
		  td[class*="mobile-header-padding-right"] {
			width: 160px !important;
			padding-right: 0 !important;
		  }
	
		  td[class="header-lg"] {
			font-size: 24px !important;
			padding-bottom: 5px !important;
		  }
	
		  td[class="header-md"] {
			font-size: 18px !important;
			padding-bottom: 5px !important;
		  }
	
		  td[class="content-padding"] {
			padding: 5px 0 30px !important;
		  }
	
		   td[class="button"] {
			padding: 5px !important;
		  }
	
		  td[class*="free-text"] {
			padding: 10px 18px 30px !important;
		  }
	
		  td[class="info-block"] {
			display: block !important;
			width: 280px !important;
			padding-bottom: 40px !important;
		  }
	
		  td[class="info-img"],
		  img[class="info-img"] {
			width: 278px !important;
		  }
		}
	  </style>
	</head>
	
	<body bgcolor="#f7f7f7">
	<table align="center" cellpadding="0" cellspacing="0" class="container-for-gmail-android" width="100%">
	  <tr>
		<td align="left" valign="top" width="100%" style="background:repeat-x url(http://s3.amazonaws.com/swu-filepicker/4E687TRe69Ld95IDWyEg_bg_top_02.jpg) #ffffff;">
		  <center>
		  <img src="http://s3.amazonaws.com/swu-filepicker/SBb2fQPrQ5ezxmqUTgCr_transparent.png" class="force-width-gmail">
			<table cellspacing="0" cellpadding="0" width="100%" bgcolor="#ffffff" background="http://s3.amazonaws.com/swu-filepicker/4E687TRe69Ld95IDWyEg_bg_top_02.jpg" style="background-color:transparent">
			  <tr>
				<td width="100%" height="80" valign="top" style="text-align: center; vertical-align:middle;">
				<!--[if gte mso 9]>
				<v:rect xmlns:v="urn:schemas-microsoft-com:vml" fill="true" stroke="false" style="mso-width-percent:1000;height:80px; v-text-anchor:middle;">
				  <v:fill type="tile" src="http://s3.amazonaws.com/swu-filepicker/4E687TRe69Ld95IDWyEg_bg_top_02.jpg" color="#ffffff" />
				  <v:textbox inset="0,0,0,0">
				<![endif]-->
				  <center>
					<table cellpadding="0" cellspacing="0" width="600" class="w320">
					  <tr>
						<td class="pull-left mobile-header-padding-left" style="vertical-align: middle;">
						  <a href=""><img width="137" height="47" src="http://s3.amazonaws.com/swu-filepicker/0zxBZVuORSxdc9ZCqotL_logo_03.gif" alt=""></a>
						</td>
	
						<td class="pull-right mobile-header-padding-right" style="color: #4d4d4d;">
							<a href="https://github.com/aawadallak/go-docker-without-air"><img width="44" height="47" src="https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png" alt="" /></a>
						  </td>
						
					  </tr>
					</table>
				  </center>
				  <!--[if gte mso 9]>
				  </v:textbox>
				</v:rect>
				<![endif]-->
				</td>
			  </tr>
			</table>
		  </center>
		</td>
	  </tr>
	  <tr>
		<td align="center" valign="top" width="100%" style="background-color: #f7f7f7;" class="content-padding">
		  <center>
			<table cellspacing="0" cellpadding="0" width="600" class="w320">
			  <tr>
				<td class="header-lg">
				  Welcome to Golang Micro services!
				</td>
			  </tr>
			  <tr>
				<td class="free-text">  
					Thank you for joining us and being this microservice, you can access more information on our github, click below to validate your account.
				</td>
				
				<tr>
					<td class="header-sm"> 
						<br>
						Your code is: <br>` +
		code +
		`</td>
				  </tr>
	  </tr>
	  <tr>
		<td align="center" valign="top" width="100%" style="background-color: #f7f7f7; height: 100px;">
		  <center>
			<table cellspacing="0" cellpadding="0" width="600" class="w320">
			  <tr>
				<td style="padding: 25px 0 25px">
				  <strong>Awsome GoLang</strong><br />
				  E-mail Dispatcher<br />
				</td>
			  </tr>
			</table>
		  </center>
		</td>
	  </tr>
	</table>
	</body>
	</html>
	`

	return str
}
