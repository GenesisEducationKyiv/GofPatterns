<?php

declare(strict_types=1);

namespace example3;

class Client
{
    public function __construct(private MailerCreatorInterface $mailerCreator)
    {
    }

    public function sendEmail(string $email, float $rate, bool $isPersonified): void
    {
        $emailCreator = $isPersonified ? new PersonifiedEmailCreator() : new CommonEmailCreator();

        $message = new EmailMessage(
            $email,
            $emailCreator->createTitle($email)->format(),
            $emailCreator->createBody($email, $rate)->format(),
        );

        $this->mailerCreator->createMailer()->send($message);
    }
}
