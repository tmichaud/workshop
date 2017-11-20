package com.tem.turnquist.learningspringboot.learningspringbootvideo;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;

/**
 * 
 * @author tmichaud
 *
 */
@Entity
public class Image
{
	@Id
	@GeneratedValue
	private Long id;

	private String name;

	// Required for JPA
	private Image()
	{
	};

	public Image(String name)
	{
		this.name = name;
	}

	public Long getId()
	{
		return id;
	}

	public void setId(Long id)
	{
		this.id = id;
	}

	public String getName()
	{
		return name;
	}

	public void setName(String name)
	{
		this.name = name;
	}

}
