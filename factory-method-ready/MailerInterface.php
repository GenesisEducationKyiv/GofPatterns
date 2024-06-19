<?php

namespace complete2;

interface MailerInterface
{
    public function send(EmailMessage $message): void;
}
