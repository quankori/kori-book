// app.controller.ts
import { Controller } from '@nestjs/common';
import { EventPattern, MessagePattern, Payload } from '@nestjs/microservices';

@Controller()
export class AppController {
  @MessagePattern('pattern_one')
  handlePatternOne(@Payload() message: any) {
    console.log('Received in pattern_one:', message);
    return { ack: 'pattern_one processed' };
  }

  @MessagePattern('pattern_two')
  handlePatternTwo(@Payload() message: any) {
    console.log('Received in pattern_two:', message);
    return { ack: 'pattern_two processed' };
  }

  // Catch-all for any unsupported patterns
  @EventPattern()
  handleUnsupportedPattern(@Payload() message: any) {
    console.warn('Received unsupported event:', message);
    return { response: 'Unsupported event received' };
  }
}
