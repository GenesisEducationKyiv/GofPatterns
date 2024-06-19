<?php

namespace complete2;

interface MailerCreatorInterface
{
    public function createMailer(): MailerInterface;
}
