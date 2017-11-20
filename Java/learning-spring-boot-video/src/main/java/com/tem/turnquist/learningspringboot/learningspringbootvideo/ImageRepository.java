package com.tem.turnquist.learningspringboot.learningspringbootvideo;

import org.springframework.data.repository.PagingAndSortingRepository;

/**
 * 
 * @author tmichaud
 *
 */
public interface ImageRepository extends PagingAndSortingRepository<Image, Long>
{
	public Image findByName(String name); // Will parse this method and look by name. - no code required.
}
