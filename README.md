# webhook-svr

hosted at https://webhook-svr-6fe715abcbd4.herokuapp.com

/mail - endpoint for the mailhook
- send to mailhook@gorunapp.live
- example: 2026-02-08T20:25:17.522067+00:00 app[web.1]: Email Recieved: {"created_at":"2026-02-08T20:25:02Z","data":{"attachments":[],"bcc":[],"cc":[],"created_at":"2026-02-08T20:25:16.032Z","email_id":"8e14c3fd-d20f-4950-a99f-da1301fab31f","from":"peterjbishop.denver@gmail.com","message_id":"\u003c56171A1C-E6EC-441E-BE2B-5E31FC8248FE@gmail.com\u003e","subject":"Test6","to":["mailhook@gorunapp.live"]},"type":"email.received"}[GIN] 2026/02/08 - 20:25:17 | 200 |     307.369µs |   52.24.126.164 | POST     "/mail"

/hook - endpoint for the webhook 


