<?php

declare(strict_types=1);

namespace example3;

class PersonifiedTitle implements EmailTitleInterface
{
    public function __construct(private readonly string $email)
    {
    }

    public function format(): string
    {
        return "Hello, $this->email!";
    }
}
