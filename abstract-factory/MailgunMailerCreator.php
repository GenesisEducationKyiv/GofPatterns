<?php

declare(strict_types=1);

namespace example3;

class MailgunMailerCreator implements MailerCreatorInterface
{
    private const RETRY_COUNT = 2;

    public function __construct(private MailgunClient $mailgunClient, private Logger $logger)
    {
    }

    public function createMailer(): MailerInterface
    {
        return new MailgunMailer($this->mailgunClient, $this->logger, self::RETRY_COUNT);
    }
}
