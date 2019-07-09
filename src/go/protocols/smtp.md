# smtp

# Links

* [Simple Mail Transfer Protocol - Wikipedia](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol)
* [RFC 821 - Simple Mail Transfer Protocol#Page 4](https://tools.ietf.org/html/rfc821#page-4)
* [SMTP Commands Reference (covers HELO/EHLO, MAIL, RCPT, DATA, RSET, VRFY, AUTH, STARTTLS etc)](https://www.samlogic.net/articles/smtp-commands-reference.htm)
* [smtp - The Go Programming Language](https://golang.org/pkg/net/smtp/)
* [smtpd - GoDoc](https://godoc.org/github.com/bradfitz/go-smtpd/smtpd)
* [RFC 5321 - Simple Mail Transfer Protocol](https://tools.ietf.org/html/rfc5321#section-4.5.2)

Simple Mail Transfer Protocol (SMTP) is an Internet standard for electronic mail (email) transmission. First defined by RFC 821 in 1982, it was updated in 2008 with Extended SMTP additions by RFC 5321, which is the protocol in widespread use today. 

SMTP defines message transport, not the message content. 

Further implementations include FTP Mail[6] and Mail Protocol, both from 1973.[7] Development work continued throughout the 1970s, until the ARPANET transitioned into the modern Internet around 1980. Jon Postel then proposed a Mail Transfer Protocol in 1980 that began to remove the mail's reliance on FTP.[8] SMTP was published as RFC 788 in November 1981, also by Postel.

# Protocol overview

SMTP is a connection-oriented, text-based protocol in which a mail sender communicates with a mail receiver by issuing command strings and supplying necessary data over a reliable ordered data stream channel, typically a Transmission Control Protocol (TCP) connection. 

An SMTP transaction consists of three command/reply sequences:

0. Dial, Sverver response "220 smtp.example.com ESMTP Postfix"

1. `HELO` or `EHLO` (res: 250)

The client sends this command to the SMTP server to identify itself and initiate the SMTP conversation. 

2. `MAIL`, `RCPT`, `DATA` 

`MAIL`: Specifies the e-mail address of the sender. If the senders e-mail address is accepted the server will reply with a 250 OK reply code. 

`RCPT`: Specifies the e-mail address of the recipient.  res 250

`DATA`: The DATA command starts the transfer of the message contents (body text, attachments etc). 

for DATA cmd it self, res 354, for data end, res 250.

3. `QUIT`

Asks the server to close the connection. If the connection can be closed the servers replies with a 221 numerical code and then is the session closed.

# SMTP transport example

```sh
S: 220 smtp.example.com ESMTP Postfix
C: HELO relay.example.com
S: 250 smtp.example.com, I am glad to meet you
C: MAIL FROM:<bob@example.com>
S: 250 Ok
C: RCPT TO:<alice@example.com>
S: 250 Ok
C: RCPT TO:<theboss@example.com>
S: 250 Ok
C: DATA
S: 354 End data with <CR><LF>.<CR><LF>
C: From: "Bob Example" <bob@example.com>
C: To: Alice Example <alice@example.com>
C: Cc: theboss@example.com
C: Date: Tue, 15 January 2008 16:02:43 -0500
C: Subject: Test message
C: 
C: Hello Alice.
C: This is a test message with 5 header fields and 4 lines in the message body.
C: Your friend,
C: Bob
C: .
S: 250 Ok: queued as 12345
C: QUIT
S: 221 Bye
{The server closes the connection}
```

# EHLO

```sh
S: 220 smtp2.example.com ESMTP Postfix
C: EHLO bob.example.com
S: 250-smtp2.example.com Hello bob.example.org [192.0.2.201]
S: 250-SIZE 14680064
S: 250-PIPELINING
S: 250 HELP
```

# Dot Encoding




