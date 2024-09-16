import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(AppModule, {
    transport: Transport.RMQ,
    options: {
      urls: ['amqp://user:password@rabbitmq:5672'],
      queue: 'nestjs_queue',  // Ensure this matches your queue in RabbitMQ
      queueOptions: {
        durable: true,
      },
    },
  });

  await app.listen();
}

bootstrap();
