<?php

namespace complete1;

interface MailerInterface
{
    public function send(EmailMessage $message): void;
}
