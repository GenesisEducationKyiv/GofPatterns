<?php

namespace example3;

interface EmailCreatorInterface
{
    public function createTitle(string $email): EmailTitleInterface;

    public function createBody(string $email, float $rate): EmailBodyInterface;
}
