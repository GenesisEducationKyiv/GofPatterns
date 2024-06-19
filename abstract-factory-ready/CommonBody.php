<?php

declare(strict_types=1);

namespace example3;

class CommonBody implements EmailBodyInterface
{
    public function format(): string
    {
        return 'Our rate is the best of the best for you! 5 UAH to 1 USD!';
    }
}
