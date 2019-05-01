import { ClientOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';

export const grpcClientOptions: ClientOptions = {
  transport: Transport.GRPC,
  options: {
    url: 'go-backend:50051',
    package: 'cats',
    protoPath: join(__dirname, '../../../libs/api-interface/src/protobuf/cats.proto'),
  },
};
