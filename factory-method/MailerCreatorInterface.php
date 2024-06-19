<?php

namespace example2;

interface MailerCreatorInterface
{
    public function createMailer(): MailerInterface;
}
