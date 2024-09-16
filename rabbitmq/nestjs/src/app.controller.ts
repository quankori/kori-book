import { Controller } from '@nestjs/common';
import { EventPattern, MessagePattern } from '@nestjs/microservices';

@Controller()
export class AppController {
  @MessagePattern('test') // Listen for specific patterns
  handleMessage(data: Record<string, unknown>): string {
    console.log('Received message from RabbitMQ:', data);
    // Process the message
    return 'Response from NestJS to Golang'; // Return response
  }
}
