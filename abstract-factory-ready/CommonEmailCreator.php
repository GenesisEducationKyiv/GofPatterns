<?php

declare(strict_types=1);

namespace example3;

class CommonEmailCreator implements EmailCreatorInterface
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
