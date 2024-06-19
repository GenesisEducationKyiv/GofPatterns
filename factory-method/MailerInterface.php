<?php

namespace example2;

interface MailerInterface
{
    public function send(EmailMessage $message): void;
}
