<?php

declare(strict_types=1);

namespace example3;

class PersonifiedEmailCreator implements EmailCreatorInterface
{
    public function __construct(private readonly UserStorage $userStorage)
    {
    }

    public function createTitle(string $email): EmailTitleInterface
    {
        return new PersonifiedTitle($this->getUserName($email));
    }

    public function createBody(string $email, float $rate): EmailBodyInterface
    {
        return new PersonifiedBody($email, $rate);
    }

    private function getUserName(string $email): string
    {
        $user = $this->userStorage->findUserByEmail($email);

        return $user->getName();
    }
}
