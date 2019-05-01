import { Min } from 'class-validator';

export class CreateCatDto {
  name: string;

  @Min(1)
  age: number;
}
