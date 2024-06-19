<?php

namespace example3;

class CommonEmailSender implements EmailCreatorInterface
{
    public function createTitle(string $email): EmailTitleInterface
    {
        return new CommonTitle();
    }

    public function createBody(string $email, float $rate): EmailBodyInterface
    {
        return new CommonBody();
    }
}
