<?php

declare(strict_types=1);

namespace example3;

class SimpleMailerCreator implements MailerCreatorInterface
{
    public function createMailer(): MailerInterface
    {
        return new SimpleMailer();
    }
}
