<?php

namespace example3;

interface MailerCreatorInterface
{
    public function createMailer(): MailerInterface;
}
