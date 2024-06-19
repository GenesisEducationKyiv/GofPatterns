<?php

declare(strict_types=1);

namespace complete1;

class MailgunMailer implements MailerInterface
{
    public function __construct(
        private MailgunClient $mailgunClient,
        private Logger $logger,
        private int $retryCount,
    ) {
    }

    public function send(EmailMessage $message): void
    {
        $retry = 0;
        while (true) {
            try {
                $this->mailgunClient->send(
                    $message->getRecipient(),
                    $message->getTitle(),
                    $message->getMessage());
                break;
            } catch (ClientException $exception) {
                $this->logger->error();
                if (++$retry >= $this->retryCount) {
                    throw $exception;
                }
            }
        }

    }
}
