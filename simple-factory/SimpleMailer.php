<?php

class SimpleMailer implements MailerInterface
{
    public function send(EmailMessage $message): void
    {
        \mail($message->getRecipient(), $message->getTitle(), $message->getMessage());
    }
}
