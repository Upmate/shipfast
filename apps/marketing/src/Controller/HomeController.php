<?php

declare(strict_types=1);

namespace ShipFast\Marketing\Controller;

use Psr\Http\Message\ResponseInterface as Response;
use Psr\Http\Message\ServerRequestInterface as Request;

class HomeController
{
    public function index(Request $request, Response $response): Response
    {
        $html = '<html><head><title>ShipFast</title></head>';
        $html .= '<body><h1>Ship faster, ship safer</h1>';
        $html .= '<p>The modern platform for rapid product delivery.</p>';
        $html .= '</body></html>';

        $response->getBody()->write($html);
        return $response->withHeader('Content-Type', 'text/html');
    }
}
