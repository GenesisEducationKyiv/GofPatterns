<?php

declare(strict_types=1);

namespace complete2;

use example3\CommonEmailCreator;
use example3\PersonifiedEmailCreator;

class Client
{
    public function __construct(
        private readonly MailerCreatorInterface $mailerCreator,
        private readonly UserStorage $userStorage,
    ) {
    }

    public function sendEmail(string $email, float $rate, bool $isPersonified): void
    {
        $emailCreator = $isPersonified
            ? new PersonifiedEmailCreator($this->userStorage)
            : new CommonEmailCreator();

        $message = new EmailMessage(
            $email,
            $emailCreator->createTitle($email)->format(),
            $emailCreator->createBody($email, $rate)->format(),
        );

        $this->mailerCreator->createMailer()->send($message);
    }
}
