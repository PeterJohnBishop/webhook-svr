# webhook-svr

hosted at https://webhook-svr-6fe715abcbd4.herokuapp.com

/mail - endpoint for the mailhook
- send to mailhook@gorunapp.live

/hook - endpoint for the webhook 

curl -X GET "https://api.resend.com/emails/0a2700a2-91ef-4c7b-92cc-565e4e682173" \
     -H "Authorization: Bearer re_5EhRSZYi_4UuznrSfcAb77iqUDCeaFsAe" \
     -H "Content-Type: application/json"