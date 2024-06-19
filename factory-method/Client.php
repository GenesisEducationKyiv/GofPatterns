<?php

declare(strict_types=1);

namespace example2;

class Client
{
    public function sendEmail(MailerCreatorInterface $mailerCreator): void
    {
        $message = new EmailMessage(
            'customer@mail.com',
            'Hello our dear customer 💖',
            'Our rate is the best of the best for you! 5 UAH to 1 USD!'
        );

        $mailerCreator->createMailer()->send($message);
    }
}
