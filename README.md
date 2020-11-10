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
		Sender:       "sender",
		AuthEmail:    "sender@mail.com",
		AuthPassword: "senderpassword",
	}

	mconfig.SetReceivers("receiver@gmail.com","receiver2@gmail.com")
	mconfig.SetSubject("Just Say Hello")
	mconfig.SetBody("Hello, <b>have a nice day</b>")

	if err := mconfig.Send(); err != nil {
		t.Logf(err.Error())
	}
	
```
 
