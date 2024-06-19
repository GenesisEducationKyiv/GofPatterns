<?php

declare(strict_types=1);

namespace complete2;

class SimpleMailerCreator implements MailerCreatorInterface
{
    public function createMailer(): MailerInterface
    {
        return new SimpleMailer();
    }
}
