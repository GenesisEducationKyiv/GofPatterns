<?php

declare(strict_types=1);

namespace complete1;

class Client
{
    public function sendEmail(MailgunClient $mailgunClient, Logger $logger, Config $config): void
    {
        $message = new EmailMessage(
            'customer@mail.com',
            'Hello our dear customer 💖',
            'Our rate is the best of the best for you! 5 UAH to 1 USD!'
        );

        $mailerFactory = new MailerFactory($mailgunClient, $logger);
        $mailerFactory->createMailer($config->getMailerType())->send($message);
    }
}
