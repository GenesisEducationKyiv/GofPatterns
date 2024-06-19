<?php

declare(strict_types=1);

namespace example3;

class PersonifiedBody implements EmailBodyInterface
{
    public function __construct(private readonly string $email, private readonly float $rate)
    {
    }

    public function format(): string
    {
        return "Welcome to the site, $this->email! Our rate is the best of the best for you! $this->rate UAH to 1 USD!";
    }
}
