<?php

declare(strict_types=1);

namespace example3;

class CommonTitle implements EmailTitleInterface
{
    public function format(): string
    {
        return 'Hello our dear customer 💖';
    }
}
