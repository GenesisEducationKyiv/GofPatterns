<?php

interface MailerInterface
{
    public function send(EmailMessage $message): void;
}
