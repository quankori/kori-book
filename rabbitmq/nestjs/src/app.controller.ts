// app.controller.ts
import { Controller } from '@nestjs/common';
import { MessagePattern, Payload } from '@nestjs/microservices';

@Controller()
export class AppController {
  @MessagePattern('pattern_one')
  handlePatternOne(@Payload() data: any) {
    console.log('Received in pattern_one:', data);
    return { ack: 'pattern_one processed' };
  }

  @MessagePattern('pattern_two')
  handlePatternTwo(@Payload() data: any) {
    console.log('Headers:', data?.properties?.headers);  // Log headers
    console.log('Received in pattern_two:', data);
    return { ack: 'pattern_two processed' };
  }
}
