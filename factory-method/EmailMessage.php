<?php

declare(strict_types=1);

namespace example2;

class EmailMessage
{
    public function __construct(
        private readonly string $recipient,
        private readonly string $title,
        private readonly string $message
    )
    {
    }

    public function getRecipient(): string
    {
        return $this->recipient;
    }

    public function getTitle(): string
    {
        return $this->title;
    }

    public function getMessage(): string
    {
        return $this->message;
    }
}
