import { Logger, OnModuleInit, ParseIntPipe, UseGuards } from '@nestjs/common';
import { Args, Mutation, Query, Resolver } from '@nestjs/graphql';
import { Client, ClientGrpc } from '@nestjs/microservices';
// import { PubSub } from 'graphql-subscriptions';
import { Observable } from 'rxjs';
import { GqlAuthGuard } from '../auth/graphql-auth.guard';
import { grpcClientOptions } from '../grpc-client.options';
import { CreateCatDto } from './dto/create-cat.dto';
import { CatEntity } from './entities/cat.entity';

// const pubSub = new PubSub();

interface CatService {
  list(data: {}): Observable<any>;
  insert(data: CreateCatDto): Observable<any>;
  get(data: { id: number }): Observable<any>;
  delete(data: { id: number }): Observable<any>;
}

@Resolver('Cat')
export class CatsResolvers implements OnModuleInit {
  @Client(grpcClientOptions)
  private readonly client: ClientGrpc;

  private catService: CatService;

  private logger = new Logger(CatsResolvers.name);

  onModuleInit() {
    this.catService = this.client.getService<CatService>('CatService');
  }

  @Query()
  @UseGuards(GqlAuthGuard)
  getCats(): Observable<any> {
    this.logger.log('getCats called');
    return  this.catService.list({});
  }

  @Query('cat')
  @UseGuards(GqlAuthGuard)
  findOneById(
    @Args('id', ParseIntPipe)
    id: number,
  ): Observable<CatEntity> {
    this.logger.log('getCat called');
    return this.catService.get({ id });
  }

  @Mutation('createCat')
  @UseGuards(GqlAuthGuard)
  create(@Args('createCatInput') args: CreateCatDto): Observable<CatEntity> {
    this.logger.log('createCat called');
    return this.catService.insert(args);
  }

  /*
  @Subscription('catCreated')
  catCreated() {
    return pubSub.asyncIterator('catCreated');
  }
  */
}
