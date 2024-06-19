<?php

namespace example2;

class SimpleMailerCreator implements MailerCreatorInterface
{
    public function createMailer(): MailerInterface
    {
        return new SimpleMailer();
    }
}
