<?php

class MailgunMailer implements MailerInterface
{
    public function __construct(
        private MailgunClient $mailgunClient,
        private Logger $logger,
        private int $retryCount = 2,
    )  {
        $this->mailgunClient = $mailgunClient;
        $this->logger = $logger;
    }

    public function send(EmailMessage $message): void
    {
        $retry = 0;
        while (true) {
            try {
                $this->mailgunClient->send($message->getRecipient(), $message->getTitle(), $message->getMessage());
                break;
            } catch (\ClientException $exception) {
                $this->logger->error($exception->getMessage());
                if (++$retry >= $this->retryCount) {
                    throw $exception;
                }
            }
        }
    }
}
