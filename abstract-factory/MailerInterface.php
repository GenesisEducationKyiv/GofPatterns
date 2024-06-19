<?php

namespace example3;

interface MailerInterface
{
    public function send(EmailMessage $message): void;
}
