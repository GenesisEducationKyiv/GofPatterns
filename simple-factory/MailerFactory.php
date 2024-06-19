<?php

class MailerFactory
{
    public const RETRY_COUNT = 2;

    public function __construct(
        private MailgunClient $mailgunClient,
        private Logger $logger
    ) {
        $this->mailgunClient = $mailgunClient;
        $this->logger = $logger;
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
