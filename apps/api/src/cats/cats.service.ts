import { CreateCatInput } from '@lazy/api-interface';
import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CatEntity } from './entities/cat.entity';

@Injectable()
export class CatsService {
  constructor(
    @InjectRepository(CatEntity)
    private readonly catRepository: Repository<CatEntity>) {
  }

  async create(cat: CreateCatInput): Promise<CatEntity> {
    return this.catRepository.save(cat);
  }

  async findAll(): Promise<CatEntity[]> {
    return this.catRepository.find();
  }

  async findOneById(id: number): Promise<CatEntity> {
    return (await this.catRepository.find({ id }))[0];
  }
}
