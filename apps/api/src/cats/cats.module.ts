import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AuthModule } from '../auth/auth.module';
import { CatsResolvers } from './cats.resolvers';
import { CatsService } from './cats.service';
import { CatEntity } from './entities/cat.entity';

@Module({
  imports: [
    AuthModule,
    TypeOrmModule.forFeature([CatEntity]),
  ],
  providers: [CatsService, CatsResolvers],
})
export class CatsModule {}
