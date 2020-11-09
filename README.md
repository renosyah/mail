# Mail Sender Library for internal use only

- how to use : 

```

import "github.com/renosyah/mail"

```

- How to Use

```

	mconfig := mail.Config{
		Host:         "smtp.gmail.com",
		Port:         587,
		Sender:       "mail.sender",
		AuthEmail:    "mail.email",
		AuthPassword: "mail.password",
	}

	mconfig.SetReceivers("example@gmail.com")
	mconfig.SetSubject("Just Say Hello")
	mconfig.SetBody("Hello, <b>have a nice day</b>")

	if err := mconfig.Send(); err != nil {
		t.Logf(err.Error())
	}
	
```
 
