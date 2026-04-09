<?php

declare(strict_types=1);

require __DIR__ . '/../vendor/autoload.php';

use ShipFast\Marketing\Controller\HomeController;
use Slim\Factory\AppFactory;

$app = AppFactory::create();

$app->get('/', [HomeController::class, 'index']);
$app->get('/health', function ($request, $response) {
    $response->getBody()->write(json_encode(['status' => 'ok', 'service' => 'shipfast-marketing']));
    return $response->withHeader('Content-Type', 'application/json');
});

$app->run();
