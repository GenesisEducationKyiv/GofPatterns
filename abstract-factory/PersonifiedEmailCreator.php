<?php

namespace example3;

class PersonifiedEmailCreator implements EmailCreatorInterface
{
    public function createTitle(string $email): EmailTitleInterface
    {
        return new PersonifiedTitle($email);
    }

    public function createBody(string $email, float $rate): EmailBodyInterface
    {
        return new PersonifiedBody($email, $rate);
    }
}
