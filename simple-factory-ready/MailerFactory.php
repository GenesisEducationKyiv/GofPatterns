<?php

declare(strict_types=1);

namespace complete1;

class MailerFactory
{
    private const RETRY_COUNT = 2;

    public function __construct(private MailgunClient $mailgunClient, private Logger $logger)
    {
    }

    public function createMailer(string $type): MailerInterface
    {
        switch ($type) {
            case 'default':
                return new SimpleMailer();
            case 'mailgun':
                return new MailgunMailer($this->mailgunClient, $this->logger, self::RETRY_COUNT);
            default:
                throw new \Exception('Wrong mailer type passed.');
        }
    }
}
