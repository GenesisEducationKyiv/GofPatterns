<?php

declare(strict_types=1);

namespace complete2;

class SimpleMailer implements MailerInterface
{
    public function send(EmailMessage $message): void
    {
        \mail($message->getRecipient(), $message->getTitle(), $message->getMessage());
    }
}
